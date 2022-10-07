import sys
import json
import pathlib
import tempfile

import arrow
import typer
from jinja2 import Template
from playwright.sync_api import sync_playwright

SUPPORTED_FORMATS = {"html", "pdf", "txt"}
SUPPORTED_THEMES = {"handmade"}


def format_data(data: dict, phone: str, email: str) -> dict:
    # Fix Profiles
    profiles: dict = {i["network"].lower(): i for i in data["basics"]["profiles"]}
    data["basics"]["profiles"] = profiles
    phone = phone.replace("-", "")

    data["basics"]["phone"] = f"({phone[:3]}) {phone[3:6]} - {phone[6:]}"
    data["basics"]["email"] = email

    # Fix Languages
    languages: list = list(zip(*(i.values() for i in data["languages"])))
    data["zipLanguages"] = languages

    # Fix Dates
    for category in {"work", "volunteer", "education", "awards", "publications"}.intersection(
        data.keys()
    ):
        for index, item in enumerate(data[category]):
            for detail in item.keys():
                if "date" in detail.lower():
                    data[category][index][detail] = arrow.get(item[detail])

    return data


def render(data: dict, ext: str, theme: str = "handmade") -> str:
    if ext == "pdf":
        ext = "html"

    path = pathlib.Path(theme, f"template.{ext}.j2").absolute()
    template = Template(path.read_text())

    return template.render({"data": data, "ext": ext})


def main(
    input: pathlib.Path = typer.Argument(pathlib.Path("resume.json"), help="an input file"),
    output: pathlib.Path = typer.Argument(pathlib.Path("./out/"), help="an output directory"),
    themes: list[str] = typer.Option(["handmade"], help="a list of themes to apply"),
    formats: list[str] = typer.Option(
        ["html", "pdf", "txt"], help="a list of file formats to generate"
    ),
    overwrite: bool = typer.Option(False, help="overwrite existing files without asking"),
    use_name: bool = typer.Option(True, help="save files named as the name in the resume data"),
    phone: str = typer.Option("555-555-5555", help="your phone number"),
    email: str = typer.Option(
        "user@email.whom", help="an email address to add to the top of the page"
    ),
) -> None:
    if input.suffix.lower() != ".json":
        sys.exit("Input file must be a .json file")

    if not set(formats).issubset(SUPPORTED_FORMATS):
        sys.exit(f"Output formats can only be: {', '.join(SUPPORTED_FORMATS)}")

    if not set(themes).issubset(SUPPORTED_THEMES):
        sys.exit(f"Theme can only be one of: {', '.join(SUPPORTED_THEMES)}")

    with open(input, "r") as json_file:
        data = json.load(fp=json_file)

    data = format_data(data=data, phone=phone, email=email)

    if use_name:
        output_name = data["basics"]["name"]

    else:
        output_name = input.stem

    for theme in themes:
        theme_dir = pathlib.Path(output, theme)
        theme_dir.mkdir(parents=True, exist_ok=True)

        for ext in formats:
            document = render(data=data, ext=ext, theme=theme)
            output_file = pathlib.Path(theme_dir, output_name).with_suffix(f".{ext}")

            if not output_file.exists() or (output_file.exists() and overwrite):
                if ext == "pdf":
                    with tempfile.NamedTemporaryFile(mode="w", suffix=".html") as pdf_loader:
                        pdf_loader.write(document)
                        pdf_loader.seek(0)

                        with sync_playwright() as playwright:
                            browser = playwright.chromium.launch()
                            page = browser.new_page(viewport={"width": 1920, "height": 1080})
                            page.goto(f"file://{pdf_loader.name}")
                            page.pdf(path=output_file)

                else:
                    output_file.write_text(document)


if __name__ == "__main__":
    typer.run(main)

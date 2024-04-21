#!env/bin/python
# Install hooks

import os
import generate_changelog

def generate_changelog_template():
    if not os.path.isfile("CHANGELOG.md"):
        generate_changelog.main()
        print("CHANGELOG.md template added")

def configure_git_config():
    text = "\thooksPath = .githooks"

    if not os.path.exists(".git"):
        print("no .git folder found")
        return

    f = open(".git/config", "r")
    lines = f.readlines()
    lines_string = ""
    for line in lines:
        # adding config in core
        if line.strip() == "[core]":
            lines_string = line + text + "\n"
            continue
        # if configuration happened before, quit
        if line.strip().find(text.strip()) != -1:
            return
        lines_string += line
    f.close()

    f = open(".git/config", "w")
    f.write(lines_string)
    f.close()

    print("git hooks activated")

def setup_virtual_env():
    # env folder creation
    print("setting up virtual env")
    os.system("python3 -m venv env")

def main():
    setup_virtual_env()
    configure_git_config()
    generate_changelog_template()

if __name__ == "__main__":
    main()

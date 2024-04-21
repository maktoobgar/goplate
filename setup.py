import os
import sys

# Change These Lines
SERVICE_NAME = "project_name"
DESCRIPTION = "Project Description"
TARGET_DIR = f"/usr/share/{SERVICE_NAME}"

Reset = "\033[0m"
Red = "\033[31m"
Green = "\033[32m"
Yellow = "\033[33m"
Blue = "\033[34m"
Purple = "\033[35m"
Cyan = "\033[36m"
Orange = "\033[33m"
Gray = "\033[37m"
White = "\033[97m"

i = 1

# Consts
CURRENT_DIR = os.path.dirname(os.path.realpath(__file__))
SERVICE_FILE_ADDR = f"/etc/systemd/system/{SERVICE_NAME}.service"


def print_started_action(started_text: str):
    global i
    print(f"\n{i} {Green}==={Reset} {Blue}{started_text}{Reset} {Green}==={Reset}")
    i += 1


def print_failed_action(failed_text: str):
    print(f"{Orange}==={Reset} {Red}{failed_text}{Reset} {Orange}==={Reset}")
    print(
        f"\n{Orange}==={Reset} {Red}Whole Installation Process Failed{Reset} {Orange}==={Reset}"
    )
    sys.exit(1)


def print_final_success_print():
    print(f"\n{Green}=== Installation finished successfully ==={Reset}")
    print(
        f"\n{White}Start Service: \t\t{Green}systemctl start {SERVICE_NAME}.service{Reset}"
    )
    print(
        f"{White}Restart Service: \t{White}systemctl restart {SERVICE_NAME}.service{Reset}"
    )
    print(f"{White}Stop Service: \t\t{Red}systemctl stop {SERVICE_NAME}.service{Reset}")
    print(
        f"{White}Check Service Status: \t{Yellow}systemctl status {SERVICE_NAME}.service{Reset}"
    )
    print(f"\nFiles And Folders of The Service is Here: {Green}{TARGET_DIR}{Reset}")


def env_address() -> str:
    return (
        f"{CURRENT_DIR}/env.yml"
        if os.path.isfile("env.yml")
        else f"{CURRENT_DIR}/env.yaml" if os.path.isfile("env.yaml") else ""
    )


def main():
    print_started_action("Check If Golang Is Installed")
    if len(os.popen("which go").readline().strip()) == 0:
        print_failed_action("Golang Not Found")

    print_started_action("Install Requirements")
    if os.system("go mod download -x") != 0:
        print_failed_action("Failed to Install Requirements")

    print_started_action("Build")
    if os.system("go build main.go") != 0:
        print_failed_action("Failed to Build")

    print_started_action("Create Service Folder")
    if os.system(f"sudo mkdir -p {TARGET_DIR}") != 0:
        print_failed_action("Failed to Create Service Folder")

    print_started_action("Copy Host Files And Folders")
    if os.system(f"sudo cp {CURRENT_DIR}/main {env_address()} {TARGET_DIR}") != 0:
        print_failed_action("Failed to Copy Content Into Service Folder")

    print_started_action("Create The Service")
    content = ""
    with open("service.service", "r") as file:
        lines = file.readlines()
        for line in lines:
            content += line
        content = content.replace("DESC", DESCRIPTION)
        content = content.replace("EXECDIR", f"{TARGET_DIR}/main")
        content = content.replace("DIR", TARGET_DIR)

    # Create SERVICE_FILE_ADDR from the content variable
    if not os.path.exists(SERVICE_FILE_ADDR):
        os.system(f"sudo echo '{content}' | sudo tee {SERVICE_FILE_ADDR}")
    else:
        os.system(f"sudo rm -f {SERVICE_FILE_ADDR}")
        os.system(f"sudo echo '{content}' | sudo tee {SERVICE_FILE_ADDR}")

    print_started_action("Start And Activate The Service")
    os.system("sudo systemctl daemon-reload")
    os.system(f"sudo systemctl enable {SERVICE_NAME}.service")
    os.system(f"sudo systemctl start {SERVICE_NAME}.service")

    print_final_success_print()


if __name__ == "__main__":
    main()

import yaml

def GetConfig(url):
    with open(url, 'r', encoding='utf8') as file:
        configs = yaml.safe_load(file)
    return configs

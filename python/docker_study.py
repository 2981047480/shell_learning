import docker

client = docker.from_env()
print(dir(client.containers))
print(client.containers.list())
def say_hello():
    print("hello world!")


def main():
    fun_list = [
        say_hello(),
        say_hello()
    ]
    for f in fun_list:
        f


if __name__ == "__main__":
    main()
from string import Template

class gen_template(object):
    """
    A template for a string.
    """
    def __init__(self, template):
        self.template = template

    def render(self, *args):
        """
        Render the template with the given arguments.
        """
        return Template(self.template).safe_substitute(*args)

    def render_from_file(self, *args):
        with open(self.template) as f:
            content = f.read()
        with open(self.template, 'w') as f:
            f.write(Template(content).safe_substitute(*args))

    def __str__(self):
        return self.template

    def __repr__(self):
        return "Template(%r)" % self.template

    def __eq__(self, other):
        return self.template == other.template

    def __ne__(self, other):
        return self.template != other.template

    def __hash__(self):
        return hash(self.template)
    
if __name__ == "__main__":
    with open('/Users/zhaozephyr/shell_learning/shell_learning/python/template_file.yml') as fp:
        content = fp.read()
    gt = gen_template(content)
    temp_dict = {
        "message": "test",
        # "file_path": __file__,
        # "aaa": "aaa"
    }
    print(gt)
    print(gt.render(temp_dict))


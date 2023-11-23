import os
import jinja2

print(jinja2.__version__)

class gen_template():
    def __init__(self, object):
        self.object = object

    def render(self, *args):
        template = jinja2.Template(self.object)
        return template.render(*args)
    
    def render_from_path(self, *args):
        env = jinja2.Environment(loader=jinja2.FileSystemLoader(os.path.dirname(self.object)))
        template = env.get_template(os.path.split(self.object)[-1])
        return template.render(*args)


# with open("/Users/zhaozephyr/shell_learning/shell_learning/python/template_file.yml") as fp:
#     content = fp.read()
# gt = gen_template(content)
# print(gt.render({"file_path": __file__}))
# print(gt.render_from_path({"file_path": __file__}))
file_path = "/Users/zhaozephyr/shell_learning/shell_learning/python/template_file.yml"
# print()
gt = gen_template(file_path)
print(gt.render_from_path({"file_path": __file__, "message": "hello jinja2", "messages": [1,2,3,4]}))
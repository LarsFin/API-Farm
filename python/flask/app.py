"""
TODO.
"""
from flask import Flask

app = Flask(__name__)

@app.route("/")
def hello_world():
    """This is a basic start."""
    return "Hello, Docker!"

if __name__ == '__main__':
    app.run()

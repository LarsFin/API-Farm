"""
TODO.
"""
from flask import Flask

app = Flask(__name__)

@app.route("/ping")
def hello_world():
    """This is a basic start."""
    return "pong"

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)
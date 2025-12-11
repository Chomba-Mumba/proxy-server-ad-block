from flask import Flask

app = Flask(__name__)

@app.route("/")
def hello_world():
    resp = "hello world"
    return resp, 200

app.run()
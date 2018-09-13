const http = require("http");

const photoHandler = require("./handlers/photo");
const notFoundHandler = require("./handlers/not_found");

const { PORT: port } = process.env;

const router = (request, response) => {
  response.setHeader("Access-Control-Allow-Origin", "*");

  // Enable Cors
  if (request.method === "OPTIONS") {
    const headers = {
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Methods": "OPTIONS, POST, GET",
      "Access-Control-Max-Age": 300
    };

    response.writeHead(204, headers);
    response.end();
    console.info(`[OPTIONS] :: ${request.url}`);
    return;
  }

  try {
    switch (request.url) {
      case "/photo":
        console.info("[POST] :: /post");
        photoHandler(request, response);
        break;
      default:
        notFoundHandler(request, response);
        break;
    }
  } catch (err) {
    response.writeHead("500", { "Content-Type": "application/json" });
    response.end(JSON.stringify({ error: { message: err.message } }));
  }
};

const server = http.createServer(router);

server.listen(port, err => {
  if (err) {
    return console.info("something bad happened", err);
  }

  console.info(`server is listening on ${port}`);
});

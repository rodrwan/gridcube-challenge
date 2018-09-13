const grpc = require("grpc");
const { parse } = require("querystring");

const { ServiceClient } = require("../service/publisher_grpc_pb");
const { GetRequest } = require("../service/publisher_pb");

const photoHandler = async (request, response) => {
  collectRequestData(request, result => {
    const address = "publisher:8091";
    const publisher = new ServiceClient(
      address,
      grpc.credentials.createInsecure()
    );

    const getRequest = new GetRequest();
    getRequest.setUsername(result.username);
    getRequest.setPassword(result.password);
    getRequest.setSize(200);
    gerRequest.setCaption(result.caption);
    publisher.uploadPicture(getRequest);

    const reply = {
      data: ""
    };

    response.writeHead("200", { "Content-Type": "application/json" });
    return response.end(JSON.stringify(reply));
  });
};

const collectRequestData = (request, callback) => {
  const FORM_URLENCODED = "application/x-www-form-urlencoded";
  if (request.headers["content-type"] === FORM_URLENCODED) {
    let body = "";
    request.on("data", chunk => {
      body += chunk.toString();
    });
    request.on("end", () => {
      callback(parse(body));
    });
  } else {
    callback(null);
  }
};

module.exports = photoHandler;

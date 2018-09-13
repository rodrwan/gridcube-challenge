const notFoundHandler = (request, response) => {
  response.writeHead("404", { "Content-Type": "application/json" });

  return response.end(
    JSON.stringify({ error: { message: "Resource not fund" } })
  );
};

module.exports = notFoundHandler;

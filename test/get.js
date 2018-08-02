const http = require("http");

module.exports = (uri) => {
  let data = "";
  return new Promise((resolve, reject) => {
    http.get(uri, (response) => {
      const statusCode = response.statusCode;
      const contentType = response.headers['content-type'];

      if (statusCode !== 200) {
        reject(`request failed with status code ${statusCode}`)
      } else if (!/^application\/json/.test(contentType)) {
        reject(`invalid response type: ${contentType}`);
      } else {
        response.on("data", (chunk) => {
          data += chunk;
        });
        response.on("end", () => {
          resolve(JSON.parse(data));
        });
      }
    });
  });
};

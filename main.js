var fs = require("fs");

var filePath = "./moby.txt";

var words = 0;
var inWord = false;

fs.readFile(filePath, (err, data) => {
  if (err) throw err;

  for (const pair of data.entries()) {
    c = pair[1];
    if (c == 32 && inWord) {
      words++;
      inWord = false;
    }

    if ((c >= 65 && c < 91) || (c >= 97 && c < 123)) {
      inWord = true;
    }
  }

  console.log(`Word count: ${words}`);
});

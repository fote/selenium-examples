var fs = require('fs');

function writeScreenShot(data, filename) {
  var stream = fs.createWriteStream(filename);
  stream.write(new Buffer(data, 'base64'));
  stream.end();
}

describe('DuckDuckGo homepage', () => {
  it('should open', async () => {
    await browser.waitForAngularEnabled(false);
    await browser.get('http://duckduckgo.com');
	browser.takeScreenshot().then(function (png) {
	  writeScreenShot(png, 'screenshot.png');
	});
    const title = await browser.getTitle(); 
    expect(title.length > 0).toBe(true);
  });
});

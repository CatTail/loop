var fs = require('fs');
var microtime = require('microtime');
var async = require('async');
var fixture = './fixtures/file.txt';
var times = 100000;

var syncStart, syncEnd, asyncStart, asyncEnd;

syncStart = microtime.nowDouble();
for (var i=0; i<times; i++) {
    fs.readFileSync(fixture);
}
syncEnd = microtime.nowDouble();
console.log('Sync time', syncEnd - syncStart);

asyncStart = microtime.nowDouble();
async.times(times, function(n, next) {
    fs.readFile(fixture, function() {
        next();
    });
}, function() {
    asyncEnd = microtime.nowDouble();
    console.log('Async time', asyncEnd - asyncStart);
});

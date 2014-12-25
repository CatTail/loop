var fs = require('fs');
var microtime = require('microtime');
var async = require('async');
var fixture = './fixtures/file.txt';
var times = 300;

var syncStart, syncEnd, asyncStart, asyncEnd;

syncStart = microtime.nowDouble();
for (var i=0; i<times; i++) {
    fs.readFileSync(fixture);
}
syncEnd = microtime.nowDouble();
console.log('Benchmark readFileSync', times, ((syncEnd - syncStart) / times * 10e9) + 'ns/op');

asyncStart = microtime.nowDouble();
async.times(times, function(n, next) {
    fs.readFile(fixture, function() {
        next();
    });
}, function() {
    asyncEnd = microtime.nowDouble();
    console.log('Benchmark readFile', times, ((asyncEnd - asyncStart) / times * 10e9) + 'ns/op');
});

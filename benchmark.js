var fs = require('fs');
var Suite = require('benchmark').Suite;
var suite = new Suite;
var fixture = './fixtures/file.txt';

suite.add('readFileSync', function() {
    fs.readFileSync(fixture);
})
.add('readFile', {
    defer: true,
    fn: function(deferred) {
        fs.readFile(fixture, function() {
            deferred.resolve();
        });
    }
})
.on('cycle', function(event) {
  console.log(String(event.target));
})
.on('complete', function() {
  console.log('Fastest is ' + this.filter('fastest').pluck('name'));
})
.run({ 'async': true });
;

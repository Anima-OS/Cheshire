const ws = new Wampy('localhost:8080', { realm: 'cheshire@anima-os.com' });

// Hello World test message.
ws.publish('example.hello', 'Hello World!');

// Take screenshot of entire screen.
ws.call('TakeScreenshot', null, {
    onSuccess: function (dataArr, dataObj) {
        console.log(dataArr);
        console.log('RPC successfully called');
        console.log('Server time is ' + dataArr[0]);
    },
    onError: function (err, detailsObj) {
        console.log('RPC call failed with error ' + JSON.stringify(err));
    }
});

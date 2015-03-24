#!/bin/env node
var nrfuart = require('nrfuart');
console.log('starting...');

//setInterval(function() {
  nrfuart.discoverAll(function onDiscover(ble_uart) {
    // enable disconnect notifications
    //ble_uart.on('disconnect', function() {
    //  console.log('disconnected!');
    //});

    // connect and setup
    console.log('connecting');
    ble_uart.connectAndSetup(function onConnect () {
      //var writeCount = 0;

      console.log('connected to: ', ble_uart.uuid);

      ble_uart.readDeviceName(function(devName) {
        //console.log('Device name:', devName);
      });

      ble_uart.on('data', function(data) {
        date = new Date();
        temp = parseFloat(data.toString());
        result = JSON.stringify({uuid: ble_uart.uuid, date: date, temp: temp});
        
        console.log('received: ', result );
        nrfuart.stopDiscoverAll(onDiscover);
        ble_uart.disconnect(onConnect);
        process.exit(0);
      });

      var TESTPATT = 'Hello world! ' ; //+ writeCount.toString();
      ble_uart.write(TESTPATT, function() {
        console.log('data sent:', TESTPATT);
        //writeCount++;
      });
    });
  });
//}, 10000);

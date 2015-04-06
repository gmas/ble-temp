#!/bin/env node
var nrfuart = require('nrfuart');

function printVal(uuid, data) {
  date = new Date();
  temp = parseFloat(data.toString());
  result = JSON.stringify({uuid: uuid, date: date, temp: temp});
  console.log(result);
}

function onDiscover(ble_uart) {
  function onConnect () {
    console.warn('connected to: ', ble_uart.uuid);

    ble_uart.on('data', function(data) {
      printVal(ble_uart.uuid, data);
      ble_uart.disconnect(onConnect); 
      process.exit(0);
    });

    var TESTPATT = 'GET' ;
    ble_uart.write(TESTPATT);
  }

  // connect to discovered device
  ble_uart.connectAndSetup(onConnect);

  // disconnect after 5s
  setTimeout(function(){ 
    console.warn('calling Disconnect on ble_uart');
    ble_uart.disconnect(onConnect); 
  }, 7000);

}

//start discovery
nrfuart.discoverAll(onDiscover);
//stop discovery and exit after 10s
setTimeout(function(){ 
  nrfuart.stopDiscoverAll(onDiscover); 
  process.exit(1);
}, 10000);

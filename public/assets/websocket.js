'use strict';

var WS = {
    i: null,
    hostname: location.hostname,
    port: 0,
    init: function (port = 8585) {
        this.port = port;

        if(this.i === null){
            this.i = new WebSocket(`ws://${this.hostname}:${this.port}/ws`);
        }
        return this.i;
    },
    send: function (request) {
        this.i.send(request);
    }
};


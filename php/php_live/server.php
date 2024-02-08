<?php
$server = new swoole_websocket_server("0.0.0.0", 8080);

$server->on('message', function (swoole_websocket_server $server, $frame) {
	foreach($server->connection_list() as $fd) {
		$server->push($fd, $frame->data,true);
	}
});

$server->start();

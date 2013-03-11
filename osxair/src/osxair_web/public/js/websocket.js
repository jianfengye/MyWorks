var socket;

socket = new WebSocket("ws://" + window.location.hostname + ":3201");

socket.onopen = function() {
	alert("websocket on onpen")
}

socket.onmessage = function(msg) {
	alert("websocket on message:" + msg.data)
}

socket.onclose = function() {
	alert("websocket on close")
}

$("input[data-type=websocket-action]").click(function(event){
	event.preventDefault();
    var target = event.target;
    var action = $(target).attr("websocket-action");
    var params = $(target).attr("websocket-params");

    var request = {}
    request.action = action;
    request.params = params;
    socket.send(JSON.stringify(request));
})
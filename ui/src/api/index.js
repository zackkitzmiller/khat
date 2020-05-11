var socket = new WebSocket("ws://localhost:8080/ws")

let connect = cb => {
    console.log("Attempting Connection")

    socket.opopen = () => {
        console.log("Successfully Connection")
    }

    socket.onmessage = msg => {
        console.log(msg)
        cb(msg)
    }

    socket.onclose = event => {
        console.log("Socket Closed Connection: ", event)
    }

    socket.onerror = error => {
        console.log("Socket Error: ", error)
    }
}

let sendMsg = (msg, username) => {
    msg = JSON.stringify({message: msg, username: username})
    console.log("sending msg: ", msg)
    socket.send(msg)
}

export { connect, sendMsg }

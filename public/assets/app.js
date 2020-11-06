var ws = WS.init(8585);
const CORRESPONDENCE = document.querySelector("#correspondence")
const USERS = document.querySelector("#users")
ws.onmessage = function (response) {
    let result = JSON.parse(response.data);
    switch (result.type) {
        case "newMessage":
            CORRESPONDENCE.innerHTML += result.text;
            break;
        case "newUser":
            USERS.innerHTML += result.text;
            if (result.append) {
                document.querySelector("#username").innerText = result.append.name;
                document.querySelector("#userImage").setAttribute("src", result.append.image);
                document.querySelector("#user").style.visibility = "visible";
            }
            break;
        case "unregisterUser":
            document.querySelector("#" + result.text).closest('li').remove();
            break;
        default:
            console.log(result)
    }
}


document.querySelector("#send").addEventListener("click", sendMessage)
document.querySelector("#message").addEventListener("keydown", (e) => {
    if("Enter" === e.code) {
        e.preventDefault();
        sendMessage();
    }
})

function sendMessage() {
    let messageTextarea = document.querySelector("#message");
    let msg = messageTextarea.value
    if (msg.length > 0) {
        ws.send(msg.trim())
    }

    messageTextarea.value = ""
}

function IsJsonString(str) {
    try {
        JSON.parse(str);
    } catch (e) {
        return false;
    }
    return true;
}

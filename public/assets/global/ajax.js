function MkRequest(url, data, method) {
    return new Promise((resolve, reject) => {

        console.log(data)


        var xhttp = new XMLHttpRequest();
        xhttp.onreadystatechange = function () {
            if (this.readyState == 4) {
                if (this.status == 200) {
                    resolve(xhttp.responseText);
                } else {
                    reject(xhttp.responseText);
                }
            }
        };
        xhttp.open(method, url, true);
        xhttp.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
        xhttp.send(JSON.stringify( data ));
    });
    // var xmlhttp = new XMLHttpRequest();   // new HttpRequest instance
    // xmlhttp.open("GET", "https://jsonplaceholder.typicode.com/posts");
    // xmlhttp.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    // xmlhttp.send(JSON.stringify({ email: "hello@user.com", response: { name: "Tester" } }));
    // return RPromise

    // if (this.readyState == 4 && this.status == 200) {
}

// MkRequest("http://localhost/", undefined, "GET")
//     .then((data) => {
//         console.log(data);
//     })
//     .catch((err) => {
//         console.error("FUCKY WUCKY", err);
//     });

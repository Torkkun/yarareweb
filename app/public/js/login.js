const URL = "http://localhost:8080/authenticate";
async function post() {
    const namebox = document.getElementById("name");
    const passbox = document.getElementById("pass");
    let data = {
        "name":namebox.value,
        "pass":passbox.value
    }
    //console.log(namebox.values,passbox.values)
    const parameter = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    }
    const result = await fetch(URL,parameter)
    .then((response) => {
        return response.json();
    })
    .catch((error) => {
        console.error(error);
    });
    alert(result);

}
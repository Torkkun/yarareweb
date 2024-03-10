const textbox = document.getElementById("search");
const URL = "http://localhost:8080/data/search?search=";
const parameter = {
    method: 'GET',
    headers: {
        'Content-Type': 'application/json'
    }
}
const table = document.getElementById("tarta");
async function get() {
    const result = await fetch(URL+textbox.value, parameter).then((respose) => {
        return respose.json();
    });
    //
    clearTbody();

    // データ行追加
    result.forEach(list => {
        let newRow = table.insertRow();
        Object.keys(list).forEach(key => {
            let newCell = newRow.insertCell();
            let newText = document.createTextNode(list[key]);
            newCell.appendChild(newText);
            //console.log(`${key} ${list[key]}`)
        })
    })
}

function clearTbody() {
    var tbodies = document.getElementsByTagName("tbody");
    for (var i = 0; i < tbodies.length; i++) {
        while (tbodies[i].rows.length > 0) {
            tbodies[i].deleteRow(0);
        }
    }
}
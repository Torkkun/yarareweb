function post() {
    var div = document.createElement('div');
    div.className = 'post';
    var list_element = document.getElementById('thread');
    let child_nodes_count = list_element.childElementCount;
    //console.log(child_nodes_count);
    div.id = child_nodes_count + 1;
    chidiv = document.getElementById('thread').appendChild(div);

    //meta
    var chidiv = document.createElement('div');
	chidiv.className = 'meta';
    chidiv.innerHTML = `
        <span class="number">${div.id}</span>
        <span class="name">
            <b>guest</b>
        </span>
        <span class="date">20XX/06/06(月) 10:10:10</span>
        <span class="uid">ID:0000</span>
    `
    document.getElementById(div.id).appendChild(chidiv);
    
    //message
    const messagebox = document.getElementById("comment");
    chidiv = document.createElement('div');
	chidiv.className = 'message';
    chidiv.innerHTML = messagebox.value;
    document.getElementById(div.id).appendChild(chidiv);
}

function dele() {
    location.reload();
}

function exploitman() {
    var div = document.createElement('div');
    div.className = 'post';
    var list_element = document.getElementById('thread');
    let child_nodes_count = list_element.childElementCount;
    //console.log(child_nodes_count);
    div.id = child_nodes_count + 1;
    chidiv = document.getElementById('thread').appendChild(div);

    //meta
    var chidiv = document.createElement('div');
	chidiv.className = 'meta';
    chidiv.innerHTML = `
        <span class="number">${div.id}</span>
        <span class="name">
            <b>guest</b>
        </span>
        <span class="date">20XX/06/06(月) 10:10:10</span>
        <span class="uid">ID:0000</span>
    `
    document.getElementById(div.id).appendChild(chidiv);
    
    //message
    chidiv = document.createElement('div');
	chidiv.className = 'message';
    chidiv.innerHTML = `<img src="1" onerror="alert('あなたのパソコンにウイルスが発見されました。OKでソフトをダウンロードできます。'); window.location = 'https://www.google.com/';">`;
    document.getElementById(div.id).appendChild(chidiv);
}
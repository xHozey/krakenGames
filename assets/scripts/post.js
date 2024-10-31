const Submit = async () => {
    const text = document.getElementById("title").value;
    const image = document.getElementById("image").files[0];
    const content = document.getElementById("content").value;
    const screens = document.getElementById("screens").files;
    const os = document.getElementById('os').value
    const processor = document.getElementById('processor').value
    const memory = document.getElementById('memory').value
    const graphics = document.getElementById('graphics').value
    const directX = document.getElementById('directX').value
    const storage = document.getElementById('storage').value
    const genre = document.getElementById('genre').value
    const developer = document.getElementById('developer').value
    const platform = document.getElementById('platform').value
    const gameSize = document.getElementById('gameSize').value
    const released = document.getElementById('released').value
    const version = document.getElementById('version').value
    let data = new FormData();
    data.append("system[os]", os);
    data.append("system[processor]", processor);
    data.append("system[memory]", memory);
    data.append("system[graphics]", graphics);
    data.append("system[direct_x]", directX);
    data.append("system[storage]", storage);
    data.append("info[genre]", genre)
    data.append("info[developer]", developer)
    data.append("info[platform]", platform)
    data.append("info[gameSize]", gameSize)
    data.append("info[released]", released)
    data.append("info[version]", version)
    data.append("title", text);
    data.append("image", image);
    data.append("content", content);
    for (let i = 0; i < screens.length; i++) {
      data.append("screens", screens[i]);
    }
    try {
      await fetch("/post", {
        method: "POST",
        body: data,
      });
    } catch (err) {
      console.error(err);
    }
    window.location.replace("http://localhost:8080/");
  };
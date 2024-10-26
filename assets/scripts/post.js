const Submit = () => {
    const text = document.getElementById("title").value;
    const image = document.getElementById("image").files[0];
    const content = document.getElementById("content").value;
    const screens = document.getElementById("screens").files;
  
    const formData = new FormData();
    formData.append("title", text);
    formData.append("image", image);
    formData.append("content", content);

    for (let i = 0; i < screens.length; i++) {
      formData.append("screens", screens[i]);
    }
  
    fetch("/post", {
      method: "POST",
      body: formData,
    })
      .then((response) => response.json())
      .then((data) => console.log(data))
      .catch((error) => console.error("Error:", error));
  };
  
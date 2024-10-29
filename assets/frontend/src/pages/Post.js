import "../index.css";

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

const PostForm = () => {
  return (
    <div className="post-form-container">
      <h2 className="form-title">Add New Game</h2>
      <div className="form-group">
        <label htmlFor="title">Title:</label>
        <input type="text" id="title" className="form-input" />
      </div>
      <div className="form-group">
        <label htmlFor="image">Image:</label>
        <input type="file" id="image" className="form-input" />
      </div>
      <div className="form-group">
        <label htmlFor="content">Content:</label>
        <textarea id="content" className="form-input form-textarea"></textarea>
      </div>
      <div className="form-group">
        <label htmlFor="screens">Screenshots:</label>
        <input type="file" multiple id="screens" className="form-input" />
      </div>
      <div className="form-section">
        <h3 className="section-title">System Requirements</h3>
        <div className="form-group">
          <label htmlFor="os">OS:</label>
          <input type="text" id="os" className="form-input" />
        </div>
        <div className="form-group">
          <label htmlFor="processor">Processor:</label>
          <input type="text" id="processor" className="form-input" />
        </div>
        <div className="form-group">
          <label htmlFor="memory">Memory:</label>
          <input type="text" id="memory" className="form-input" />
        </div>
        <div className="form-group">
          <label htmlFor="graphics">Graphics:</label>
          <input type="text" id="graphics" className="form-input" />
        </div>
        <div className="form-group">
          <label htmlFor="directX">DirectX:</label>
          <input type="text" id="directX" className="form-input" />
        </div>
        <div className="form-group">
          <label htmlFor="storage">Storage:</label>
          <input type="text" id="storage" className="form-input" />
        </div>
      </div>
      <div className="form-section">
        <h3 className="section-title">Game Information</h3>
        <div className="form-group">
          <label htmlFor="genre">Genre:</label>
          <input type="text" id="genre" className="form-input" />
        </div>
        <div className="form-group">
          <label htmlFor="developer">Developer:</label>
          <input type="text" id="developer" className="form-input" />
        </div>
        <div className="form-group">
          <label htmlFor="platform">Platform:</label>
          <input type="text" id="platform" className="form-input" />
        </div>
        <div className="form-group">
          <label htmlFor="gameSize">Game Size:</label>
          <input type="text" id="gameSize" className="form-input" />
        </div>
        <div className="form-group">
          <label htmlFor="releasedDate">Released by:</label>
          <input type="text" id="released" className="form-input" />
        </div>
        <div className="form-group">
          <label htmlFor="version">Version:</label>
          <input type="text" id="version" className="form-input" />
        </div>
      </div>
      <button onClick={Submit} className="submit-button">
        Submit
      </button>
    </div>
  );
};

export default PostForm;

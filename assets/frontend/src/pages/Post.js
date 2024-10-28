import "../index.css";

const Submit = () => {
    
}


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
            </div>Game
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
            <h3 className="section-title"> Information</h3>
            <div className="form-group">
              <label htmlFor="genre">Genre (comma-separated):</label>
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
              <label htmlFor="releasedDate">Released Date:</label>
              <input type="date" id="releasedDate" className="form-input" />
            </div>
            <div className="form-group">
              <label htmlFor="version">Version:</label>
              <input type="text" id="version" className="form-input" />
            </div>
          </div>
          <button onClick={Submit} className="submit-button">Submit</button>
      </div>
    );
  };

export default PostForm;

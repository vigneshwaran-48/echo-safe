

document.addEventListener("DOMContentLoaded", function() {
  document.getElementById("sidebar-toggle").addEventListener("click", function() {
    document.getElementById("sidebar").classList.toggle("w-0");
    document.getElementById("sidebar").classList.toggle("w-[200px]");
  });

  document.querySelectorAll("editor > .note-content-input").forEach(element => {
    element.addEventListener("keyup", e => {
      const mdToHTMLContent = marked.parse(e.target.value);
      document.querySelector(`#${e.target.dataset.preview} > .content`).innerHTML = DOMPurify.sanitize(mdToHTMLContent);
    })
  });

  document.querySelectorAll(".note-title-input").forEach(element => {
    element.addEventListener("keyup", e => {
      document.querySelector(`#${e.target.dataset.preview} > .note-title`).innerText = e.target.value;
    })
  });

  document.addEventListener("keyup", e => {
    if (e.altKey && e.key === "p") {
      e.preventDefault()
      const activeEditor = document.querySelector(".editor.active"); // Handling for multiple editor window with .active class
      if (!activeEditor) {
        console.log("No active editor exists");
        return;
      }

      // Parsing here for cases where the note content id not edited.
      const activeEditorContent = document.querySelector(".editor.active > .note-content-input");
      const mdToHTMLContent = marked.parse(activeEditorContent.value);
      document.querySelector(`#${activeEditorContent.dataset.preview} > .content`).innerHTML = DOMPurify.sanitize(mdToHTMLContent);

      activeEditor.classList.toggle("hidden");
      document.getElementById(activeEditor.dataset.preview).classList.toggle("hidden");
    }
  });
});

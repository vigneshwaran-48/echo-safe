

document.addEventListener("DOMContentLoaded", function() {
  document.getElementById("sidebar-toggle").addEventListener("click", function() {
    document.getElementById("sidebar").classList.toggle("w-0");
    document.getElementById("sidebar").classList.toggle("w-[200px]");
  });

  Array.from(document.getElementsByClassName("editor")).forEach(element => {
    element.addEventListener("keyup", e => {
      const mdToHTMLContent = marked.parse(e.target.value);
      document.getElementById(e.target.dataset.preview).innerHTML = DOMPurify.sanitize(mdToHTMLContent);
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
      activeEditor.classList.toggle("hidden");
      document.getElementById(activeEditor.dataset.preview).classList.toggle("hidden");
    }
  })
});

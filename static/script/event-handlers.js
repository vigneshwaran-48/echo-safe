
const getPosition = (e) => {
  let posx = 0;
  let posy = 0;
  if (e.pageX || e.pageY) {
    posx = e.pageX;
    posy = e.pageY;
  } else if (e.clientX || e.clientY) {
    posx = e.clientX + document.body.scrollLeft +
      document.documentElement.scrollLeft;
    posy = e.clientY + document.body.scrollTop +
      document.documentElement.scrollTop;
  }
  return {
    x: posx,
    y: posy
  }
}

document.addEventListener("DOMContentLoaded", function() {

  window.onNoteSidebarContextClick = (navLink, e) => {
    e.preventDefault();
    e.stopPropagation();
    const menu = document.querySelector("#sidebar-context-menu");
    const pos = getPosition(e);
    menu.style.left = `${pos.x}px`;
    menu.style.top = `${pos.y}px`;
    menu.classList.add("scale-100");
    document.querySelector("#note-context-menu-delete").dataset.id = navLink.dataset.id;
  }

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

  document.body.addEventListener("oncreatenote", e => {
    document.querySelectorAll(".nav-link").forEach(navLink => {
      navLink.classList.remove("text-primary-text", "bg-dark-hover");
    });
    document.querySelector(`#note-sidenav-${e.detail.id}`).classList.add("text-primary-text", "bg-dark-hover");
    history.pushState(null, "", `/notes/${e.detail.id}`);
    document.title = e.detail.title;
  })

  document.querySelector("#note-context-menu-delete").addEventListener("click", e => {
    document.querySelector(`#note-sidenav-${e.target.dataset.id} > .delete`).click();
  })

  // Outside click listeners
  document.body.addEventListener("click", () => {
    document.querySelector("#sidebar-context-menu").classList.remove("scale-100");
  })
});

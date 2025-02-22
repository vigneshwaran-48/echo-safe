

document.addEventListener("DOMContentLoaded", function() {
  document.getElementById("sidebar-toggle").addEventListener("click", function() {
    document.getElementById("sidebar").classList.toggle("w-0");
    document.getElementById("sidebar").classList.toggle("w-[200px]");
  });
});

let dataProject = [];

function addProject(event) {
  event.preventDefault();
  let projectName = document.getElementById("input-project-name").value;
  let startDate = document.getElementById("input-start-date").value;
  let endDate =  document.getElementById("input-end-date").value;
  let description = document.getElementById("input-description").value;
  let image = document.getElementById("input-imeg").files;

  // durasi project

  let mulai = new Date(startDate);
  let akhir = new Date(endDate);
  let selisih = akhir.getTime() - mulai.getTime();
  let durasi = Math.ceil(selisih / (1000 * 3600 * 24 * 30));

  // icon

  const nodejsIcon = '<i class="fa-brands fa-node-js"></i>';
  const nextjsIcon = '<i class="fa-brands fa-google-play"></i>';
  const reactjsIcon = '<i class="fa-brands fa-react"></i>';
  const typescriptIcon = '<i class="fa-brands fa-java"></i>';

  // icon ke checkbox

  let checkboxNodejs = document.getElementById("node-js").checked
    ? nodejsIcon
    : "";
  let checkboxNextjs = document.getElementById("next-js").checked
    ? nextjsIcon
    : "";
  let checkboxReactjs = document.getElementById("react-js").checked
    ? reactjsIcon
    : "";
  let checkboxTypescript = document.getElementById("type-script").checked
    ? typescriptIcon
    : "";

  // url gambar yang di input
  image = URL.createObjectURL(image[0]);
  console.log(image);

  // menampung data project

  let project = {
    projectName,
    startDate,
    endDate,
    durasi,
    description,
    image,
    checkboxNodejs,
    checkboxNextjs,
    checkboxReactjs,
    checkboxTypescript,
  };

  dataProject.push(project);
  console.log(dataProject);

  renderProject();
}

function renderProject() {
  document.getElementById("project").innerHTML = "";

  for (let index = 0; index < dataProject.length; index++) {
    document.getElementById("project").innerHTML += `
    <div class="grid-item1">
          <img src="${dataProject[index].image}" alt="gambar-hehe" />
          <div class="judul-project-deadline">
            <a href="my-project-content.html"><p class="judul-project">${dataProject[index].projectName}</p></a>
            <p class="deadline-project">Duration: ${dataProject[index].durasi} bulan</p>
          </div>
          <div class="project-detail">
            <p>
              ${dataProject[index].description}
            </p>
          </div>
          <div class="font-awesome-grid">
            ${dataProject[index].checkboxNodejs}
            ${dataProject[index].checkboxNextjs}
            ${dataProject[index].checkboxReactjs}
            ${dataProject[index].checkboxTypescript}
          </div>
          <div class="btn-group-post-project">
            <button>Edit</button>
            <button>Read</button>
          </div>
        </div>
        `;
  }
}

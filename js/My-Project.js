let dataProject = [];

function addProject(event) {
  event.preventDefault();
  let projectName = document.getElementById("input-project-name").value;
  let startDate = document.getElementById("input-start-date").value;
  let endDate = document.getElementById("input-end-date").value;
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
    postAt: new Date(),
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
            <a href="#"><p class="judul-project">${
              dataProject[index].projectName
            }</p></a>
            <p class="deadline-project">Duration: ${
              dataProject[index].durasi
            } bulan</p>
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
            <p style="font-size: 15px; color: grey; letter-spacing: 0px;">${getDistanceTime(
              dataProject[index].postAt
            )} | ${getFullTime(dataProject[index].postAt)}
            </p>
          </div>
          <div class="btn-group-post-project">
            <button>Edit</button>
            <button>Read</button>
          </div>
        </div>
        `;
  }
}
// cara mencari waktu saat post
function getFullTime(time) {
  // console.log("get full time");

  // full time
  // let time = new Date();
  // console.log(time);

  // mencari waktu spesifik dd-mm-yy hours:minutes
  let monthName = [
    "Jan",
    "Feb",
    "Mar",
    "Apr",
    "May",
    "Jun",
    "Jul",
    "Aug",
    "Sep",
    "Oct",
    "Des",
  ];

  let date = time.getDate();
  let monthIndex = time.getMonth();
  let year = time.getFullYear();
  let hours = time.getHours();
  let minutes = time.getMinutes();

  if (hours <= 9) {
    hours = "0" + hours;
  }
  if (minutes <= 9) {
    minutes = "0" + minutes;
  }

  return `${date}-${monthName[monthIndex]}-${year} ${hours}:${minutes} WIB`;
}

// mencari jarak antara waktu post dan waktu sebenarnya
function getDistanceTime(time) {
  let timeNow = new Date();
  let timePost = time;

  let distance = timeNow - timePost;

  let milisecond = 1000; // 1detik in ms
  let secondInHours = 3600; // 1jam in second
  let hoursInDay = 24; // 1day in hours

  let distanceDay = Math.floor(
    distance / (milisecond * secondInHours * hoursInDay)
  );
  let distanceHours = Math.floor(distance / (milisecond * 60 * 60));
  let distanceMinutes = Math.floor(distance / (milisecond * 60));
  let distanceSecond = Math.floor(distance / milisecond);

  if (distanceDay > 0) {
    return `${distanceDay} Day Ago`;
  } else if (distanceHours > 0) {
    return `${distanceHours} Hours Ago`;
  } else if (distanceMinutes > 0) {
    return `${distanceMinutes} Minutes Ago`;
  } else {
    return `${distanceSecond} Seconds Ago`;
  }
}

// set real time distancetime
// setInterval(function () {
//   renderProject();
// }, 2000);

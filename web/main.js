function load() {
    getData();
}

function getData() {
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() { 
        if (this.readyState == 4 && this.status == 200){
            var data = JSON.parse(this.responseText);
            if (data.length > 0){
               fillTable(data);
            }
        }
    };
    xhttp.open("GET", "http://localhost:8070/getmodels", true);
    xhttp.setRequestHeader("Content-Type", "application/json");
    xhttp.send();
}

function fillTable(data) {
    var trow = "";
    console.log(data);
    data.forEach(element => {
        trow += "<tr>";
        trow += "<td>" + element.name + "</td>";
        trow += "<td>" + element.set + "</td>";
        trow += "<td>" + element.year + "</td>";
        trow += "<td>" + element.modelnumber + "</td>";
        trow += "<td>" + element.manufacturer + "</td>";
        trow += "<td> <div class=\"btn-group pull-right\"> <button type=\"button\"class=\"btn btn-info\" onlick=\"edit()\" id=\"edit\"> Edit </button> <button type=\"button\" class=\"btn btn-danger\"onlick=\"remove(this)\" id=\"remove\">Remove</button></div></td></tr>";
    })
    document.getElementById('data').innerHTML = trow;
}

function addData(){
    var name = document.getElementById("name").value;
    var set = document.getElementById("set").value;
    var year = document.getElementById("year").value;
    var mfact = document.getElementById("mfact").value;
    var modelNumber = document.getElementById("model").value;
    var photographed = document.getElementById("photo").value;

    if (name.length == 0 || modelnumber.length == 0){
        alert("You have to enter Model name and Model number")
    } else { 
        var xhttp = new XMLHttpRequest();
        xhttp.onreadystatechange = function() {
            if (this.readyState == 4 && this.status == 200){
                getData();
            }
        };
        xhttp.open("POST", "http://localhost:8070/addmodel/", true);
        xhttp.setRequestHeader("Content-Type", "application/json");

        var data = {
            name: name, 
            set: set,
            year: year,
            manufacturer: mfact,
            modelnumber: modelNumber
        };
        var str = JSON.stringify(data);
        console.log(str);
        xhttp.send(str);
    }
}

function remove(object){
    if (typeof(object) == "object") {
        var curentRow = ($object).closest("tr");
        var name = curentRow.find("td:eq(0)").text();
        var modelnum = curentRow.find("td:eq(3)").text();
        var xhttp = new XMLHttpRequest();
        xhttp.onreadystatechange = function() {
            if (this.readyState == 4 && this.status == 200){
                getData();
            }
        };
        xhttp.open("DELETE", "http://localhost:8070/removemodel/"+name+"/"+modelnum, true);
        xhttp.setRequestHeader("Content-Type", "application/json");
        xhttp.send();
    } else {
        return false;
    }
}


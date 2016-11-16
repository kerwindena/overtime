function loadDayRange(begin, end) {
  var xobj = new XMLHttpRequest();
  xobj.overrideMimeType("application/json");
  xobj.open('GET', '/api/getDateRange/?begin=' + begin + '&end=' + end, true);
  xobj.onreadystatechange = function () {
    if (xobj.readyState == 4 && xobj.status == "200") {
      // Required use of an anonymous callback as .open will NOT return a value but simply returns undefined in asynchronous mode
      loadDayRangeCallback(JSON.parse(xobj.responseText));
    }
  };
  xobj.send(null);
}

function loadDayRangeCallback(value) {
  var table = document.querySelector("table tbody")
  var rowExample = document.querySelector("table tbody tr.rowExample")
  value.forEach(function(rowData) {
    if(!document.querySelector("table tbody tr[data-day='"+rowData["Day"]+"']")) {
      var row = rowExample.cloneNode(true);
      row.classList.remove("rowExample")

      row.setAttribute("data-day", rowData["Day"])

      row.querySelector(".day").textContent = rowData["Day"]
      row.querySelector(".date").textContent = rowData["Date"]

      var rows = document.querySelectorAll("table tbody tr[data-day]");
      if(rows.length === 0 || rowData["Day"] - 1 == rows[rows.length-1].getAttribute("data-day")) {
        table.append(row)
      } else {
        inserted = false;
        rows.forEach(function(r) {1
          if(rowData["Day"] - 1 == r.getAttribute("data-day")) {
            table.insertBefore(row, r.nextSibling);
            inserted = true;
          }
        });
        if(!inserted) {
          table.insertBefore(row, rows[0]);
        }
      }
    }
  });
}

function loadPreviousDays() {
  var rows = document.querySelectorAll("table tbody tr[data-day]");
  var firstRowDay = 1*rows[0].getAttribute("data-day");
  loadDayRange(firstRowDay - 2, firstRowDay - 1);
}

function loadSubsequentDays() {
  var rows = document.querySelectorAll("table tbody tr[data-day]");
  var lastRowDay = 1*rows[rows.length-1].getAttribute("data-day");
  loadDayRange(lastRowDay + 1, lastRowDay + 2);
}


loadDayRange(17250, 17251)

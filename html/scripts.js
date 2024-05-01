$(document).ready( function () {
    $('#myTable').DataTable();
});

let table = new DataTable('#myTable', {
    ajax: {
        url: 'http://localhost:8080/api/json?get=sample',
        dataSrc: ''},
    columns: [
        {data: 'id', visible: false},
        {data: 'name'},
        {data: 'age'},
        {data: 'email'},
        {data: 'phone'},
    ],
});

table.on('click', 'tbody tr', function () {
    let data = table.row(this).data();
     alert('You clicked on id: ' + data.id);
});

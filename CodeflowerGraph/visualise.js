
var mysql = require('mysql');
var con = mysql.createConnection({
    host: "localhost",
    user: "Enter Username..",
    password: "Enter Password Here..",
    database: "extractdb"
});

con.connect(function(err) {
    if (err) throw err;
    con.query("SELECT name, size, language FROM repo", function (err1, result1, fields1) {
        if (err1) throw err1;
        var obj1 = result1;
        var json1 = JSON.stringify(obj1); 
        var fs1 = require('fs');
        fs1.writeFile('repoData.json', json1, 'utf8', function(err){
            if(err){
                return console.log(err);
            }
        });
    }); 
    con.query("SELECT * FROM org", function (err2, result2, fields2) {
        if (err2) throw err2;
        var obj2 = result2;
        var json2 = JSON.stringify(obj2);
//        var fs2 = require('fs');
/*        fs2.writeFile('orgs.json', json2, 'utf8', function(err){
            if(err){
                return console.log(err);
            }
        });*/
    });
    con.query("SELECT * FROM follower", function (err3, result3, fields3) {
        if (err3) throw err3;
        var obj3 = result3;
        var json3 = JSON.stringify(obj3);
    
        var fs3 = require('fs');
        fs3.writeFile('followers.json', json3, 'utf8', function(err){
            if(err){
                return console.log(err);
            }
        });
    });
  });
var x1 >= 0;
var x2 >= 0;

minimize z: 9*x1 + 15*x2 ;

s.t. st1: 9*x1 + 2*x2 >= 54 ;
s.t. st2: x1 + 5*x2 >= 25 ;
s.t. st3: x1 + x2 >= 13 ;

end ;

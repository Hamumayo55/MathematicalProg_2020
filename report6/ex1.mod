var x1 >= 0;
var x2 >= 0;

minimize z: -3*x1 + -2*x2 ;

s.t. st1: 2*x1 + 5*x2 <= 40 ;
s.t. st2: 3*x1 + x2 <= 30 ;
s.t. st3: 3*x1 + 4*x2 <= 39 ;

end ;

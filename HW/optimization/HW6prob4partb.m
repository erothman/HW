c = [23;1;-16;-1;52;-6;-12];
b = [62; 1861/32; 3];
A=[-6,-5,25,3,-85,4,30;24,-2,28,6,-55,1,-9;9,-5,11,2,-55,-1,19];
[xsol, objective, basis] = simplextwoEricRothman(A,b,c);
xsol
objective
basis
c = [10;10;0;-3;-1;-3];
b = [26;17;1];
A = [-1,2,3,6,9,8;-2,3,1,1,6,8;1,1,-1,1,1,3];
[x,y,z] = simplextwoEricRothman(A,b,c);
x2 = sprintf('%f ', x);
z2 = sprintf('%d ', z);
fprintf("The solution vector is: [%s]\n", x2);
fprintf("The optimal objection value is %f \n", y);
fprintf("The basis for the optimal solution is: [%s]\n", z2);
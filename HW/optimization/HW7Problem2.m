c = [5;1;7;3;0;1;2;3];
b = [4;2;7];
A=[6,2,-5,7,9,-3,1,9;5,-6,5,4,1,9,8,-1;1,7,3,-1,-1,4,0,8];
I = -eye(3);
Aslack = [A,I];
zero = transpose(zeros(1,3));
cslack = [c;zero];
ctranspose = transpose(cslack);
Tableau = [1,-ctranspose,0;zero,-Aslack,-b]

fprintf("We will use the dual simplex method starting with the basis of the slack variables.\n");
fprintf("First we need to decide which row to pivot on.");
fprintf("The fourth row has the smallest b value, so we shall pivot there.\n");
fprintf("Then we need to decide the column to pivot on using the ratio test.\n");
fprintf("We will only look at the columns with a negative value in that row.")
Tableau(1,2)/Tableau(4,2)
Tableau(1,3)/Tableau(4,3)
Tableau(1,4)/Tableau(4,4)
Tableau(1,7)/Tableau(4,7)
Tableau(1,9)/Tableau(4,9)

fprintf("The column with the minimum ratio is column 3 so we shall pivot on row 4 column 3.\n")

Tableau(4,:) = Tableau(4,:)/Tableau(4,3)
Tableau(1,:) = Tableau(1,:) - Tableau(1,3)*Tableau(4,:)
Tableau(2,:) = Tableau(2,:) - Tableau(2,3)*Tableau(4,:)
Tableau(3,:) = Tableau(3,:) - Tableau(3,3)*Tableau(4,:)

fprintf("Since atleast one b value is less than 0, be pivot again.\n");
fprintf("The row with the smallest b value is row 3 with a value of -8,");
fprintf(" so we shall pivot on that row.\n");
fprintf("To decide the column we will do the ratio test again only looking at the negative values in the row.\n");
Tableau(1,2)/Tableau(3,2)
Tableau(1,4)/Tableau(3,4)
Tableau(1,5)/Tableau(3,5)
Tableau(1,6)/Tableau(3,6)
Tableau(1,7)/Tableau(3,7)
Tableau(1,8)/Tableau(3,8)
Tableau(1,9)/Tableau(3,9)

fprintf("The column with the minimum ratio is column 7 so we shall pivot on row 3 column 7.\n");

Tableau(3,:) = Tableau(3,:)/Tableau(3,7)
Tableau(1,:) = Tableau(1,:) - Tableau(1,7)*Tableau(3,:)
Tableau(2,:) = Tableau(2,:) - Tableau(2,7)*Tableau(3,:)
Tableau(4,:) = Tableau(4,:) - Tableau(4,7)*Tableau(3,:)

fprintf("Since atleast one b value is less than 0, be pivot again.\n");
fprintf("The row with the smallest b value is row 2 with a value of -4.6666,");
fprintf(" so we shall pivot on that row.\n");
fprintf("To decide the column we will do the ratio test again only looking at the negative values in the row.\n");
Tableau(1,2)/Tableau(2,2)
Tableau(1,5)/Tableau(2,5)
Tableau(1,6)/Tableau(2,6)
Tableau(1,8)/Tableau(2,8)
Tableau(1,9)/Tableau(2,9)

fprintf("The column with the minimum ratio is column 6 so we shall pivot on row 2 column 6.\n");
Tableau(2,:) = Tableau(2,:)/Tableau(2,6)
Tableau(1,:) = Tableau(1,:) - Tableau(1,6)*Tableau(2,:)
Tableau(3,:) = Tableau(3,:) - Tableau(3,6)*Tableau(2,:)
Tableau(4,:) = Tableau(4,:) - Tableau(4,6)*Tableau(2,:)

fprintf("Now since all b values are nonnegative and all c values are nonpositive, we have arrived at optimum.\n");

fprintf("The x values that are the optimal solution to x can be read straight from the simplex.");

x_values = [0;0.7069;0;0;0.5;0.6379;0;0;0;0;0];

fprintf("The optimal solution is");
x_values
fprintf("of the form:\n[x1\n x2\n x4\n x4\n x5\n x6\n x7\n x8\n s1\n s2\n s3]\n");
fprintf("where the s are the slack variables.\n");

fprintf("Optimal objective value is:");
Tableau(1,13)

fprintf("To find the y values we need to solve for B inverse where B is the basis for the solution");

b=[2,9,-3;-6,1,9;7,-1,4]
bcalc = [b,eye(3)]
bcalc(1,:) = bcalc(1,:)/bcalc(1,1)
bcalc(2,:) = bcalc(2,:) - bcalc(2,1)*bcalc(1,:)
bcalc(3,:) = bcalc(3,:) - bcalc(3,1)*bcalc(1,:)
bcalc(2,:) = bcalc(2,:)/bcalc(2,2)
bcalc(1,:) = bcalc(1,:) - bcalc(1,2)*bcalc(2,:)
bcalc(3,:) = bcalc(3,:) - bcalc(3,2)*bcalc(2,:)
bcalc(3,:) = bcalc(3,:)/bcalc(3,3)
bcalc(2,:) = bcalc(2,:) - bcalc(2,3)*bcalc(3,:)
bcalc(1,:) = bcalc(1,:) - bcalc(1,3)*bcalc(3,:)

fprintf("So B inverse is:")
Binverse = [bcalc(1,4),bcalc(1,5),bcalc(1,6);bcalc(2,4),bcalc(2,5),bcalc(2,6);bcalc(3,4),bcalc(3,5),bcalc(3,6)]

CbT = transpose([c(2);c(5);c(6)])


fprintf("The optimal solution is");
y = transpose(mtimes(CbT, Binverse))
fprintf("of the form:\n[y1\n y2\n y3]\n");

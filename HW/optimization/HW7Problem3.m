c = [5;2;1;0;1;3;2];
b = [6;3];
A=[-1,3,4,6,-2,4,1;8,-2,5,-3,0,1,5];
I = -eye(2);
Aslack = [A,I];
zero = transpose(zeros(1,2));
cslack = [c;zero];
ctranspose = transpose(cslack);
Tableau = [1,-ctranspose,0;zero,-Aslack,-b]

fprintf("We will use the dual simplex method starting with the basis of the slack variables.\n");
fprintf("First we need to decide which row to pivot on.");
fprintf("The second row has the smallest b value, so we shall pivot there.\n");
fprintf("Then we need to decide the column to pivot on using the ratio test.\n");
fprintf("We will only look at the columns with a negative value in that row.")
Tableau(1,3)/Tableau(2,3)
Tableau(1,4)/Tableau(2,4)
Tableau(1,5)/Tableau(2,5)
Tableau(1,7)/Tableau(2,7)
Tableau(1,8)/Tableau(2,8)
fprintf("The column with the minimum ratio is column 5 so we shall pivot on row 2 column 5.\n");
fprintf("Note that the column has a ratio of 0 and that the column has a 0 at the top so there will be degeneracy.\n");
fprintf("That means the objective value will not change.");

Tableau(2,:) = Tableau(2,:)/Tableau(2,5)
Tableau(1,:) = Tableau(1,:) - Tableau(1,5)*Tableau(2,:)
Tableau(3,:) = Tableau(3,:) - Tableau(3,5)*Tableau(2,:)

fprintf("Since atleast one b value is less than 0, be pivot again.\n");
fprintf("The row with the smallest b value is row 3 with a value of -6,");
fprintf(" so we shall pivot on that row.\n");
fprintf("To decide the column we will do the ratio test again only looking at the negative values in the row.\n");
Tableau(1,2)/Tableau(3,2)
Tableau(1,4)/Tableau(3,4)
Tableau(1,7)/Tableau(3,7)
Tableau(1,8)/Tableau(3,8)

fprintf("The column with the minimum ratio is column 4 so we shall pivot on row 3 column 4.\n");
Tableau(3,:) = Tableau(3,:)/Tableau(3,4)
Tableau(1,:) = Tableau(1,:) - Tableau(1,4)*Tableau(3,:)
Tableau(2,:) = Tableau(2,:) - Tableau(2,4)*Tableau(3,:)

fprintf("Now since all b values are nonnegative and all c values are nonpositive, we have arrived at optimum.\n");

fprintf("The x values that are the optimal solution to x can be read straight from the simplex.");

x_values = [0;0;0.8571;0.4286;0;0;0;0;0];

fprintf("The optimal solution is");
x_values
fprintf("of the form:\n[x1\n x2\n x4\n x4\n x5\n x6\n x7\n s1\n s2]\n");
fprintf("where the s are the slack variables.\n");

fprintf("Optimal objective value is:");
Tableau(1,11)

fprintf("To find the y values we need to solve for B inverse where B is the basis for the solution");

b=[4,6;5,-3]
bcalc = [b,eye(2)]
bcalc(1,:) = bcalc(1,:)/bcalc(1,1)
bcalc(2,:) = bcalc(2,:) - bcalc(2,1)*bcalc(1,:)
bcalc(2,:) = bcalc(2,:)/bcalc(2,2)
bcalc(1,:) = bcalc(1,:) - bcalc(1,2)*bcalc(2,:)

fprintf("So B inverse is:")
Binverse = [bcalc(1,3),bcalc(1,4);bcalc(2,3),bcalc(2,4)]

CbT = transpose([c(3);c(4)])


fprintf("The optimal solution is");
y = transpose(mtimes(CbT, Binverse))
fprintf("of the form:\n[y1\n y2]\n");
c = [1;2;3;4];
b = [3;4];
A=[1,-2,-2,-2;-2,1,-2,-2];
I = -eye(2);
Aslack = [A,I];
zero = transpose(zeros(1,2));
cslack = [c;zero];
ctranspose = transpose(cslack);
Tableau = [1,-ctranspose,0;zero,-Aslack,-b]

fprintf("We will use the dual simplex method starting with the basis of the slack variables.\n");
fprintf("First we need to decide which row to pivot on.");
fprintf("The third row has the smallest b value, so we shall pivot there.\n");
fprintf("Then we need to decide the column to pivot on using the ratio test.\n");
fprintf("We will only look at the columns with a negative value in that row.");
fprintf("There is only one negative value in that row, so we will pivot on the second column where that value is.\n")

Tableau(2,:) = Tableau(2,:)/Tableau(2,2)
Tableau(1,:) = Tableau(1,:) - Tableau(1,2)*Tableau(2,:)
Tableau(3,:) = Tableau(3,:) - Tableau(3,2)*Tableau(2,:)

fprintf("Since atleast one b value is less than 0, be pivot again.\n");
fprintf("The row with the smallest b value is row 3 with a value of -10,");
fprintf("Here is where a problem arises though.\n");
fprintf("Normally you would use the ratio test to decide which columnt to pivot on but ");
fprintf("every non b value in the third row is non negative, so no matter how much we raise ")
fprintf("any column, the c values will never go above 0.\n")
fprintf("So we can raise the dual variables by any amount and it won't violate the conditions.\n")
fprintf("So the problem is unbounded in the dual and infeasible in the LP.\n")

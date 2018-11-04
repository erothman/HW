C = [1,7,-37,94,-9,76,-146]
b = [26;18;34]
A = [7,7,45,-1,3,-53,-68;9,-5,27,-115,7,-129,42;5,-3,63,-96,10,-109,86]
rightSide = [0;b]
leftSide = [1;0;0;0]
middle = [-C; A]
Tableau = [leftSide,middle,rightSide]

fprintf("This is the pretableau\n")
fprintf("The starting Basis we will work with is columns 1,2,5\n")
format short
Tableau(2,:) = Tableau(2,:)/(Tableau(2,2))
Tableau(1,:) = Tableau(1,:)-(Tableau(2,:)*Tableau(1,2))
Tableau(3,:) = Tableau(3,:)-(Tableau(2,:)*Tableau(3,2))
Tableau(4,:) = Tableau(4,:)-(Tableau(2,:)*Tableau(4,2))
fprintf("Reduce the first column to rref\n")
Tableau(3,:) = Tableau(3,:)/(Tableau(3,3))
Tableau(1,:) = Tableau(1,:)-(Tableau(3,:)*Tableau(1,3))
Tableau(2,:) = Tableau(2,:)-(Tableau(3,:)*Tableau(2,3))
Tableau(4,:) = Tableau(4,:)-(Tableau(3,:)*Tableau(4,3))
fprintf("Reduce the 2nd column to rref\n")
Tableau(4,:) = Tableau(4,:)/(Tableau(4,6))
Tableau(1,:) = Tableau(1,:)-(Tableau(4,:)*Tableau(1,6))
Tableau(3,:) = Tableau(3,:)-(Tableau(4,:)*Tableau(3,6))
Tableau(2,:) = Tableau(2,:)-(Tableau(4,:)*Tableau(2,6))
fprintf("Reduce the 5th column to rref\n")
fprintf("No column would cause a change if we pivot on it. But we can ")
fprintf("still pivot on column 7 despite it no changing the obj. func. value.\n")
fprintf("We are doing this to see if there is just a stall, or if it is already the best solution.\n")
fprintf("The only row in that column with a non-negative number is row 4.\n")
fprintf("So we will pivot on row 4.\n")
fprintf("The two ratios to decide the pivot are:")
Tableau(4,:) = Tableau(4,:)/(Tableau(4,8))
Tableau(1,:) = Tableau(1,:)-(Tableau(4,:)*Tableau(1,8))
Tableau(2,:) = Tableau(2,:)-(Tableau(4,:)*Tableau(2,8))
Tableau(3,:) = Tableau(3,:)-(Tableau(4,:)*Tableau(3,8))
fprintf("A similar thing happens again. So this time lets pivot on column 5.\n")
fprintf("All entries in the column are positive, except the top one, so we need ")
fprintf("to compare all ratios:")
Tableau(2,9)/Tableau(2,6)
Tableau(3,9)/Tableau(3,6)
Tableau(4,9)/Tableau(4,6)
fprintf("Both the top and bottom rows have the same ration.\n")
fprintf("Since if we pivot on the bottom row, we will end up at the beginning ")
fprintf("again, with a basis of cols 1,2,5, we will pivot on the first row.\n")
Tableau(2,:) = Tableau(2,:)/(Tableau(2,6))
Tableau(1,:) = Tableau(1,:)-(Tableau(2,:)*Tableau(1,6))
Tableau(4,:) = Tableau(4,:)-(Tableau(2,:)*Tableau(4,6))
Tableau(3,:) = Tableau(3,:)-(Tableau(2,:)*Tableau(3,6))
fprintf("Once again we get a change of 0 when pivoting on the only column ")
fprintf("that is not part of the basis or has a non-negative chage.\n")
fprintf("The only value on column 1 we can pivot on is the first row, which ")
fprintf("will just undo our last step.\nSo all three answers are the same.\n")
fprintf("So in this case, lets let the first answer, the original basis, be ")
fprintf("the submitted answer.\n")
fprintf("So the min value for the objective function is -22 which occurs at ")
fprintf("the values of:\n")
fprintf("[x1    [0\n x2     2\n x3     0\n x4  =  0\n x5     4\n x6     0\n x7]    0]\n")

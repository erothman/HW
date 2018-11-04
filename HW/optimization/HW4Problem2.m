C = [-4,63,7,-2,-2,0,21]
b = [107;55;25]
A = [8,-226,-33,10,9,49,-1;9,-199,-51,10,3,25,-25;2,24,45,-6,3,-45,-15]
rightSide = [0;b]
leftSide = [1;0;0;0]
middle = [-C; A]
Tableau = [leftSide,middle,rightSide]

fprintf("This is the pretableau\n")
fprintf("The starting Basis we will work with is columns 1,3,4\n")
format short
Tableau(2,:) = Tableau(2,:)/(Tableau(2,2))
Tableau(1,:) = Tableau(1,:)-(Tableau(2,:)*Tableau(1,2))
Tableau(3,:) = Tableau(3,:)-(Tableau(2,:)*Tableau(3,2))
Tableau(4,:) = Tableau(4,:)-(Tableau(2,:)*Tableau(4,2))
fprintf("Reduce the first column to rref\n")
Tableau(3,:) = Tableau(3,:)/(Tableau(3,4))
Tableau(1,:) = Tableau(1,:)-(Tableau(3,:)*Tableau(1,4))
Tableau(2,:) = Tableau(2,:)-(Tableau(3,:)*Tableau(2,4))
Tableau(4,:) = Tableau(4,:)-(Tableau(3,:)*Tableau(4,4))
fprintf("Reduce the 3rd column to rref\n")
Tableau(4,:) = Tableau(4,:)/(Tableau(4,5))
Tableau(1,:) = Tableau(1,:)-(Tableau(4,:)*Tableau(1,5))
Tableau(3,:) = Tableau(3,:)-(Tableau(4,:)*Tableau(3,5))
Tableau(2,:) = Tableau(2,:)-(Tableau(4,:)*Tableau(2,5))
fprintf("Reduce the 4th column to rref\n")
fprintf("The Column that would cause the largest change is column 6:")
fprintf(" it is the one with the largest number in the top row.\n")
fprintf("So we will pivot based on that column.\n")
fprintf("We can ignore pivoting on the 1st row since that is negative.")
fprintf("The two ratios to decide the pivot are:")
Tableau(3,9)/Tableau(3,7)
Tableau(4,9)/Tableau(4,7)
fprintf("Since the ratio when pivoted on the 4th row is smaller we shall ")
fprintf("pivot there because that would be the row that becomes 0 first.\n")
Tableau(4,:) = Tableau(4,:)/(Tableau(4,7))
Tableau(1,:) = Tableau(1,:)-(Tableau(4,:)*Tableau(1,7))
Tableau(2,:) = Tableau(2,:)-(Tableau(4,:)*Tableau(2,7))
Tableau(3,:) = Tableau(3,:)-(Tableau(4,:)*Tableau(3,7))
fprintf("The only pivitable column with a positive number at the top is ")
fprintf("column 2.\nThe thing is though that all the numbers in that ")
fprintf("column that can be pivoted on are negative.\nSo that means no ")
fprintf("matter how much we increase x2, none of the conditions will ever ")
fprintf("reach 0, so no condition will ever be violated.\nThat means the ")
fprintf("linear program is unbounded and as such does not have an answer.\n")
fprintf("There is no way to minimize the linear program given.\n")
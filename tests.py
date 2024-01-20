from hstest.stage_test import *
from hstest.test_case import TestCase

CheckResult.correct = lambda: CheckResult(True, '')
CheckResult.wrong = lambda feedback: CheckResult(False, feedback)


class CalcTest(StageTest):
    on_exit = False

    def generate(self) -> List[TestCase]:
        return [TestCase(stdin=['/help', self.test_1_1, self.test_1_2]),
                TestCase(stdin=['5 + 1 + 4', self.test_2_1, self.test_2_2, self.test_2_3,
                                self.test_2_4, self.test_2_5]),
                TestCase(stdin=['n = 5', self.test_3_1, self.test_3_2, self.test_3_3, self.test_3_4, self.test_3_5,
                                self.test_3_6, self.test_3_7, self.test_3_8, self.test_3_9,
                                self.test_3_10, self.test_3_11]),
                TestCase(stdin=['a = 9\nb=2\nc = 1', self.test_4_1, self.test_4_2, self.test_4_3, self.test_4_4,
                                self.test_4_5, self.test_4_6]),
                # new test cases for uppercase variables
                TestCase(stdin=['a  =  3\nb= 4\nc =5', self.test_new_1, self.test_new_2, self.test_new_3,
                                self.test_new_4, self.test_new_5, self.test_new_6, self.test_new_7, self.test_new_8]),
                TestCase(stdin=['var1 = 1', self.test_5_1, self.test_5_2, self.test_5_3, self.test_5_4,
                                self.test_5_5, self.test_5_6])]

    # test of functionality of previous steps #################################
    # test of /help command
    def test_1_1(self, output):
        output = str(output).lower().strip()
        if len(output.split(" ")) < 3:
            return CheckResult.wrong("It seems like there wasn't any \"help\" message.")
        return ""

    # input empty string
    def test_1_2(self, output):
        output = str(output)
        if len(output) != 0:
            return CheckResult.wrong("Incorrect response to an empty string. "
                                     "The program should not print anything.")
        self.on_exit = True
        return "/exit"

    # sum of positive numbers
    def test_2_1(self, output):
        output = str(output).lower().strip()
        if output != "10":
            return CheckResult.wrong("The program cannot sum more than two numbers.")
        return "23 - 17 - 13"

    # sum of numbers with different signs & negative answer
    def test_2_2(self, output):
        output = str(output).lower().strip()
        if output != "-7":
            return CheckResult.wrong("Incorrect sum of positive and negative numbers.")
        return "33 + 20 + 11 + 49 - 32 - 9 + 1 - 80 + 4"

    # testing a big amount of numbers
    def test_2_3(self, output):
        output = str(output).lower().strip()
        if output != "-3":
            return CheckResult.wrong("The program cannot process a large number of numbers.")
        return "11 - 7 - 4"

    # deleted due to excessive complexity
    # def test_6(self, output):
    #     output = str(output).lower().strip()
    #     if output != "0":
    #         return CheckResult.wrong("The problem when sum is equal to 0 has occurred.")
    #     return "5 --- 2 ++++++ 4 -- 2 ---- 1" output = 10

    # the sum is zero
    def test_2_4(self, output):
        output = str(output).lower().strip()
        if output != "0":
            return CheckResult.wrong("The program cannot process multiple operations with several operators.")
        return "/start"

    # test of nonexistent command
    def test_2_5(self, output):
        output = str(output).lower().strip()
        if "unknown" not in output:
            return CheckResult.wrong("The program should print \"Unknown command\" " +
                                     "when a nonexistent command is entered.")
        self.on_exit = True
        return "/exit"

    # testing basic assignments
    def test_3_1(self, output):
        if len(output) != 0:
            return CheckResult.wrong("Unexpected reaction after assignment."
                                     "The program should not print anything in this case.")
        return "m=2"

    # assignment without spaces
    def test_3_2(self, output):
        if len(output) != 0:
            return CheckResult.wrong("Unexpected reaction after assignment."
                                     "The program should not print anything in this case.")
        return "a    =  7"

    # assignment with extra spaces
    def test_3_3(self, output):
        if len(output) != 0:
            return CheckResult.wrong("Unexpected reaction after assignment."
                                     "The program should not print anything in this case.")
        return " c=  a "

    # assign the value of another variable
    def test_3_4(self, output):
        if len(output) != 0:
            return CheckResult.wrong("Unexpected reaction after assignment."
                                     "The program should not print anything in this case.")
        return "a"

    # test printing values of the variables
    def test_3_5(self, output):
        output = str(output).lower().strip()
        if output != "7":
            return CheckResult.wrong("The variable stores not a correct value.")
        return "c"

    def test_3_6(self, output):
        output = str(output).lower().strip()
        if output != "7":
            return CheckResult.wrong("The variable stores not a correct value.")
        return "test = 0"

    # trying to assign a new variable after other operations
    def test_3_7(self, output):
        if len(output) != 0:
            return CheckResult.wrong("Unexpected reaction after assignment."
                                     "The program should not print anything in this case.")
        return "test"

    # checking if the assignment was successful
    def test_3_8(self, output):
        output = str(output).lower().strip()
        if output != "0":
            return CheckResult.wrong("The variable stores not a correct value.")
        return "test = 1"

    # trying to reassign
    def test_3_9(self, output):
        if len(output) != 0:
            return CheckResult.wrong("Unexpected reaction after assignment."
                                     "The program should not print anything in this case.")
        return "a = test"

    # trying to reassign with the value of another variable
    def test_3_10(self, output):
        if len(output.strip()) != 0:
            return CheckResult.wrong("Unexpected reaction after assignment."
                                     "The program should not print anything in this case.")
        return "a"

    def test_3_11(self, output):
        output = str(output).lower().strip()
        if output != "1":
            return CheckResult.wrong("The program could not reassign variable.")
        self.on_exit = True
        return "/exit"

    # testing operations with variables
    # creating some variables (a = 9, b=2 and c=1)
    # and testing simple adding
    def test_4_1(self, output):
        if len(output) != 0:
            return CheckResult.wrong("Unexpected reaction after assignment."
                                     "The program should not print anything in this case.")
        return "a + b"

    def test_4_2(self, output):
        output = str(output).lower().strip()
        if output != "11":
            return CheckResult.wrong("The program cannot perform operations with variables. " +
                                     "For example, addition operation.")
        return "b - a"

    # subtracting test
    def test_4_3(self, output):
        output = str(output).lower().strip()
        if output != "-7":
            return CheckResult.wrong("The program cannot perform operations with variables. " +
                                     "For example, subtraction operation.")
        return "b + c"

    # adding a negative number
    def test_4_4(self, output):
        output = str(output).lower().strip()
        if output != "3":
            return CheckResult.wrong("The program cannot perform operations with variables. " +
                                     "For example, addition operation.")
        return "b - c"

    # subtracting a negative number
    def test_4_5(self, output):
        output = str(output).lower().strip()
        if output != "1":
            return CheckResult.wrong("The program cannot perform operations with variables. " +
                                     "For example, subtraction operation.")
        return "a -- b - c + 3 --- a ++ 1"

    # testing multiple operations
    def test_4_6(self, output):
        output = str(output).lower().strip()
        if output != "5":
            return CheckResult.wrong("The program cannot perform several operations in one line.")
        self.on_exit = True
        return "/exit"

    # a set of negative test ##################################################
    # testing invalid variable name
    def test_5_1(self, output):
        output = str(output).lower().strip()
        if "invalid" not in output:
            return CheckResult.wrong("The name of a variable should contain only Latin letters.")
        return "var = 2a"

    # testing invalid value
    def test_5_2(self, output):
        output = str(output).lower().strip()
        if "invalid" not in output:
            return CheckResult.wrong("The value can be an integer number or a value of another variable.")
        return "c = 7 -  1 = 5"

    # test multiple equalization
    def test_5_3(self, output):
        output = str(output).lower().strip()
        if "invalid" not in output:
            return CheckResult.wrong("The program could not handle a invalid assignment.")
        return "variable = f"

    # testing assignment by unassigned variable
    def test_5_4(self, output):
        output = str(output).lower().strip()
        if "unknown" not in output and "invalid" not in output:
            return CheckResult.wrong("The program should not allow an assignment by unassigned variable.")
        return "variable = 777"

    # checking case sensitivity
    def test_5_5(self, output):
        if len(output) != 0:
            return CheckResult.wrong("Unexpected reaction after assignment. "
                                     "The program should not print anything in this case.")
        return "Variable"

    def test_5_6(self, output):
        output = str(output).lower().strip()
        if "unknown" not in output:
            return CheckResult.wrong("The program should be case sensitive.")
        self.on_exit = True
        return "/exit"

    # new test cases for uppercase variables
    def test_new_1(self, output):
        if len(output) != 0:
            return CheckResult.wrong("Unexpected reaction after assignment."
                                     "The program should not print anything in this case.")
        return "a + b - c"

    def test_new_2(self, output):
        output = str(output).lower().strip()
        if output != "2":
            return CheckResult.wrong("The program cannot perform operations with variables. " +
                                     "For example, addition or subtraction operation.")
        return "b - c + 4 - a"

    def test_new_3(self, output):
        output = str(output).lower().strip()
        if output != "0":
            return CheckResult.wrong("The program cannot perform operations with variables. " +
                                     "For example, addition or subtraction operation.")
        return "X = 800\nY=4\nZ= 5"

    def test_new_4(self, output):
        if len(output) != 0:
            return CheckResult.wrong("Unexpected reaction after assignment with uppercase variables."
                                     "The program should not print anything in this case.")
        return "X + Y + Z"

    def test_new_5(self, output):
        output = str(output).lower().strip()
        if output != "809":
            return CheckResult.wrong("The program cannot perform operations with uppercase variables. " +
                                     "For example, addition operation.")
        return "BIG = 9000"

    def test_new_6(self, output):
        output = str(output).lower().strip()
        if len(output) != 0:
            return CheckResult.wrong("Unexpected reaction after assignment with uppercase variables. "
                                     "The program should not print anything in this case.")
        return "BIG"

    def test_new_7(self, output):
        output = str(output).lower().strip()
        if output != "9000":
            return CheckResult.wrong("The variable stores not a correct value.")
        return "big"

    def test_new_8(self, output):
        output = str(output).lower().strip()
        if "unknown" not in output:
            return CheckResult.wrong("The program should be case sensitive.")
        self.on_exit = True
        return "/exit"

    def check(self, reply: str, attach) -> CheckResult:
        if self.on_exit:
            reply = reply.strip().lower().split('\n')
            self.on_exit = False
            if 'bye' not in reply[-1]:
                return CheckResult.wrong("Your program didn't print \"bye\" after entering \"/exit\".")
            else:
                return CheckResult.correct()
        else:
            return CheckResult.wrong("The program ended prematurely")


if __name__ == '__main__':
    CalcTest("calculator.calculator").run_tests()

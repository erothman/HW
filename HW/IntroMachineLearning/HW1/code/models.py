# Modified and completed by Eric Rothman
# email: erothma6@jhu.edu
import numpy as np
import math
from scipy.sparse import csr_matrix


class Model(object):

    def __init__(self):
        self.num_input_features = None

    def fit(self, X, y):
        """ Fit the model.

        Args:
            X: A compressed sparse row matrix of floats with shape
                [num_examples, num_features].
            y: A dense array of ints with shape [num_examples].
        """
        raise NotImplementedError()

    def predict(self, X):
        """ Predict.

        Args:
            X: A compressed sparse row matrix of floats with shape
                [num_examples, num_features].

        Returns:
            A dense array of ints with shape [num_examples].
        """
        raise NotImplementedError()


class Useless(Model):

    def __init__(self):
        super().__init__()
        self.reference_example = None
        self.reference_label = None

    def fit(self, X, y):
        self.num_input_features = X.shape[1]
        # Designate the first training example as the 'reference' example
        # It's shape is [1, num_features]
        self.reference_example = X[0, :]
        # Designate the first training label as the 'reference' label
        self.reference_label = y[0]
        self.opposite_label = 1 - self.reference_label

    def predict(self, X):
        if self.num_input_features is None:
            raise Exception('fit must be called before predict.')
        # Perhaps fewer features are seen at test time than train time, in
        # which case X.shape[1] < self.num_input_features. If this is the case,
        # we can simply 'grow' the rows of X with zeros. (The copy isn't
        # necessary here; it's just a simple way to avoid modifying the
        # argument X.)
        num_examples, num_input_features = X.shape
        if num_input_features < self.num_input_features:
            X = X.copy()
            X._shape = (num_examples, self.num_input_features)
        # Or perhaps more features are seen at test time, in which case we will
        # simply ignore them.
        if num_input_features > self.num_input_features:
            X = X[:, :self.num_input_features]
        # Compute the dot products between the reference example and X examples
        # The element-wise multiply relies on broadcasting; here, it's as if we first
        # replicate the reference example over rows to form a [num_examples, num_input_features]
        # array, but it's done more efficiently. This forms a [num_examples, num_input_features]
        # sparse matrix, which we then sum over axis 1.
        dot_products = X.multiply(self.reference_example).sum(axis=1)
        # dot_products is now a [num_examples, 1] dense matrix. We'll turn it into a
        # 1-D array with shape [num_examples], to be consistent with our desired predictions.
        dot_products = np.asarray(dot_products).flatten()
        # If positive, return the same label; otherwise return the opposite label.
        same_label_mask = dot_products >= 0
        opposite_label_mask = ~same_label_mask
        y_hat = np.empty([num_examples], dtype=np.int)
        y_hat[same_label_mask] = self.reference_label
        y_hat[opposite_label_mask] = self.opposite_label
        return y_hat


class Perceptron(Model):

    def __init__(self, learn, times):
        super().__init__()
        self.learning_rate = learn
        self.iterations = times
        pass

    def fit(self, X, y):
        self.w = np.zeros((1,X.shape[1]))
        self.num_input_features = X.shape[1]
        k = 0
        Xarray = X.toarray()
        while k < self.iterations:
            for i, xi in enumerate(Xarray):
                total = np.multiply(xi, self.w).sum(axis=1)
                predict = 1 if total[0] >= 0.0 else -1
                trueVal = -1 if y[i] == 0 else 1
                self.w = np.add(self.w, xi*self.learning_rate*(trueVal - predict))
            k += 1
        pass

    def predict(self, X):
        if self.num_input_features is None:
            raise Exception('fit must be called before predict.')
        # Perhaps fewer features are seen at test time than train time, in
        # which case X.shape[1] < self.num_input_features. If this is the case,
        # we can simply 'grow' the rows of X with zeros. (The copy isn't
        # necessary here; it's just a simple way to avoid modifying the
        # argument X.)
        num_examples, num_input_features = X.shape
        if num_input_features < self.num_input_features:
            X = X.copy()
            X._shape = (num_examples, self.num_input_features)
        # Or perhaps more features are seen at test time, in which case we will
        # simply ignore them.
        if num_input_features > self.num_input_features:
            X = X[:, :self.num_input_features]

        Xarray = X.toarray()

        dot_products = np.multiply(self.w, Xarray).sum(axis=1)

        same_label_mask = dot_products >= 0
        opposite_label_mask = ~same_label_mask
        y_hat = np.empty([num_examples], dtype=np.int)
        y_hat[same_label_mask] = 1
        y_hat[opposite_label_mask] = 0
        return y_hat
        pass


class Logistic(Model):

    def __init__(self, online_learning_rate, online_training_iterations):
        super().__init__()
        self.learning_rate = online_learning_rate
        self.iterations = online_training_iterations
        pass

    def _g(self, z):
        if z < 0:
            return 1 - (1/(1 + np.exp(z)))
        return 1/(1 + np.exp(-1 * z))

    def fit(self, X, y):
        self.w = np.zeros((X.shape[1]))
        self.num_input_features = X.shape[1]
        k = 0
        Xarray = X.toarray()
        while k < self.iterations:
            for i, xi in enumerate(Xarray):
                dot_product = np.dot(self.w, xi)
                delta = y[i]*self._g(-1*dot_product)*xi - (1-y[i])*self._g(dot_product)*xi
                self.w = self.w + self.learning_rate * delta
            k += 1
        pass

    def predict(self, X):
        if self.num_input_features is None:
            raise Exception('fit must be called before predict.')
        # Perhaps fewer features are seen at test time than train time, in
        # which case X.shape[1] < self.num_input_features. If this is the case,
        # we can simply 'grow' the rows of X with zeros. (The copy isn't
        # necessary here; it's just a simple way to avoid modifying the
        # argument X.)
        num_examples, num_input_features = X.shape
        if num_input_features < self.num_input_features:
            X = X.copy()
            X._shape = (num_examples, self.num_input_features)
        # Or perhaps more features are seen at test time, in which case we will
        # simply ignore them.
        if num_input_features > self.num_input_features:
            X = X[:, :self.num_input_features]

        Xarray = X.toarray()
        prob = np.empty([num_examples], dtype=np.float64)

        for i, xi in enumerate(Xarray):
            prob[i] = self._g(np.dot(self.w, xi))

        same_label_mask = (prob >= 0.5)
        opposite_label_mask = ~same_label_mask
        y_hat = np.empty([num_examples], dtype=np.int)
        y_hat[same_label_mask] = 1
        y_hat[opposite_label_mask] = 0
        return y_hat
        pass


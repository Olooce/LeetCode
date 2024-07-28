import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

class S47_permutationsII {
  /**
   * Generates all unique permutations of the given integer array.
   * 
   * @param nums The input array containing integers, with duplicates.
   * @return A list of lists, where each list represents a unique permutation of the array.
   */
  public List<List<Integer>> permuteUnique(int[] nums) {
    List<List<Integer>> ans = new ArrayList<>(); // To store the final list of unique permutations.
    Arrays.sort(nums); // Sort the array to facilitate the identification of duplicates.
    dfs(nums, new boolean[nums.length], new ArrayList<>(), ans); // Start depth-first search.
    return ans;
  }

  /**
   * A helper method that performs depth-first search to generate permutations.
   *
   * @param nums The sorted input array.
   * @param used A boolean array indicating whether an element is used in the current permutation.
   * @param path The current permutation being constructed.
   * @param ans The list to store all unique permutations.
   */
  private void dfs(int[] nums, boolean[] used, List<Integer> path, List<List<Integer>> ans) {
    // Base case: If the current path length equals the nums array length,
    // a complete permutation has been formed.
    if (path.size() == nums.length) {
      ans.add(new ArrayList<>(path)); // Add the current permutation to the answer list.
      return; // Backtrack.
    }

    // Iterate through the array to construct permutations.
    for (int i = 0; i < nums.length; ++i) {
      if (used[i]) // Skip if the element at index i is already used in the current path.
        continue;

      // Skip if the current element is a duplicate of the previous element,
      // and the previous element has not been used in the current path.
      if (i > 0 && nums[i] == nums[i - 1] && !used[i - 1])
        continue;

      used[i] = true; // Mark the current element as used.
      path.add(nums[i]); // Add the current element to the permutation path.
      dfs(nums, used, path, ans); // Recursively build the next level of the permutation.
      path.remove(path.size() - 1); // Backtrack by removing the last element from the path.
      used[i] = false; // Mark the element as unused for further exploration.
    }
  }

  /**
   * The main method to test the permuteUnique method with a larger array.
   *
   * @param args Command-line arguments (not used).
   */
  public static void main(String[] args) {
    S47_permutationsII solution = new S47_permutationsII();

    // Example input array with elements, including duplicates
    int[] nums = {1, 2, 2, 3};

    // Get all unique permutations
    List<List<Integer>> permutations = solution.permuteUnique(nums);

    // Print the permutations
    for (List<Integer> permutation : permutations) {
      System.out.println(permutation);
    }
  }
}

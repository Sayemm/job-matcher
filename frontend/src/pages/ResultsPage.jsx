import { useState, useEffect } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import { getJobsByCluster } from '../services/api';
import JobCard from '../components/JobCard';
import Pagination from '../components/Pagination';
import Loader from '../components/Loader';

export default function ResultsPage() {
  const location = useLocation();
  const navigate = useNavigate();
  const [jobs, setJobs] = useState([]);
  const [currentPage, setCurrentPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);

  const initialData = location.state?.data;

  useEffect(() => {
    if (!initialData) {
      navigate('/');
      return;
    }

    if (currentPage === 1 && initialData.jobs) {
      setJobs(initialData.jobs);
      setTotalPages(Math.ceil(initialData.total_in_cluster / 20));
    } else {
      fetchJobs(currentPage);
    }
  }, [currentPage, initialData, navigate]);

  const fetchJobs = async (page) => {
    setIsLoading(true);
    setError(null);

    try {
      const result = await getJobsByCluster(initialData.cluster_id, page, 20);
      setJobs(result.jobs || result.data);
      setTotalPages(result.total_pages || Math.ceil(result.total_in_cluster / 20));
    } catch (err) {
      setError('Failed to load jobs');
      console.error('Error fetching jobs:', err);
    } finally {
      setIsLoading(false);
    }
  };

  const handlePageChange = (page) => {
    setCurrentPage(page);
    window.scrollTo({ top: 0, behavior: 'smooth' });
  };

  const handleNewSearch = () => {
    navigate('/');
  };

  if (!initialData) {
    return null;
  }

  return (
    <div className="min-h-screen bg-gray-50">
      <div className="bg-white shadow-sm border-b border-gray-200">
        <div className="container mx-auto px-4 py-6">
          <div className="flex items-center justify-between">
            <div>
              <h1 className="text-3xl font-bold text-gray-900">
                Matching Jobs
              </h1>
              <p className="text-gray-600 mt-1">
                Found {initialData.total_in_cluster} jobs that match your profile
              </p>
            </div>
            <button
              onClick={handleNewSearch}
              className="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
            >
              New Search
            </button>
          </div>

          <div className="mt-4 flex items-center gap-4">
            <div className="bg-green-50 border border-green-200 rounded-lg px-4 py-2">
              <span className="text-sm text-green-800 font-medium">
                Match Score: {(initialData.match_score * 100).toFixed(1)}%
              </span>
            </div>
            <div className="bg-blue-50 border border-blue-200 rounded-lg px-4 py-2">
              <span className="text-sm text-blue-800 font-medium">
                Cluster: {initialData.cluster_id}
              </span>
            </div>
          </div>
        </div>
      </div>

      <div className="container mx-auto px-4 py-8">
        {error && (
          <div className="mb-6 p-4 bg-red-50 border border-red-200 rounded-lg text-red-700">
            {error}
          </div>
        )}

        {isLoading ? (
          <div className="flex items-center justify-center py-20">
            <Loader />
          </div>
        ) : (
          <>
            <div className="mb-6 text-sm text-gray-600">
              Showing {(currentPage - 1) * 20 + 1} - {Math.min(currentPage * 20, initialData.total_in_cluster)} of {initialData.total_in_cluster} jobs
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {jobs.map((job) => (
                <JobCard key={job.id} job={job} />
              ))}
            </div>

            {totalPages > 1 && (
              <Pagination
                currentPage={currentPage}
                totalPages={totalPages}
                onPageChange={handlePageChange}
              />
            )}
          </>
        )}
      </div>
    </div>
  );
}
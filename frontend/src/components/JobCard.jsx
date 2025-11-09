export default function JobCard({ job }) {
  const formatSalary = (min, max, currency = 'USD') => {
    if (!min && !max) return 'Not specified';
    
    const formatter = new Intl.NumberFormat('en-US', {
      style: 'currency',
      currency: currency || 'USD',
      minimumFractionDigits: 0,
    });

    if (min && max) {
      return `${formatter.format(min)} - ${formatter.format(max)}`;
    } else if (min) {
      return `From ${formatter.format(min)}`;
    } else {
      return `Up to ${formatter.format(max)}`;
    }
  };

  return (
    <div className="bg-white rounded-lg shadow-md hover:shadow-lg transition-shadow p-6 border border-gray-200">
      <h3 className="text-xl font-semibold text-gray-900 mb-2">
        {job.title}
      </h3>

      <p className="text-gray-600 font-medium mb-3">
        {job.company_name || 'Company not specified'}
      </p>

      <div className="flex items-center gap-4 mb-3 text-sm text-gray-600">
        <div className="flex items-center gap-1">
          <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
          </svg>
          <span>{job.location || 'Location not specified'}</span>
        </div>

        {job.remote_allowed && (
          <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
            Remote
          </span>
        )}
      </div>

      {job.experience_level && (
        <div className="mb-3">
          <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
            {job.experience_level}
          </span>
        </div>
      )}

      {(job.min_salary || job.max_salary) && (
        <div className="mb-3 text-sm">
          <span className="font-medium text-gray-700">Salary: </span>
          <span className="text-gray-600">
            {formatSalary(job.min_salary, job.max_salary)}
          </span>
        </div>
      )}

      {job.description && (
        <p className="text-gray-600 text-sm mb-4 line-clamp-3">
          {job.description}
        </p>
      )}

      {job.job_posting_url && (
        <a
          href={job.job_posting_url}
          target="_blank"
          rel="noopener noreferrer"
          className="inline-flex items-center justify-center w-full px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors"
        >
          Apply Now
          <svg className="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
          </svg>
        </a>
      )}
    </div>
  );
}